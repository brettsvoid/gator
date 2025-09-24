package commands

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/brettsvoid/gator/internal/database"
)

/**
* Usage: gator browse --limit <int> --sort <asc|desc> --filter <string>
 */
func HandlerBrowse(s *State, cmd Command, user database.User) error {
	parsed := ParseArgs(cmd.Args)

	// defaults
	limit := 2
	sortOrder := "desc" // newest first by default
	filterQ := ""

	// --limit or -l
	if v, ok := parsed.Flags["limit"]; ok {
		n, err := strconv.Atoi(v)
		if err != nil || n <= 0 {
			return fmt.Errorf("invalid --limit: %v", err)
		}
		limit = n
	}

	// --sort asc|desc -s
	if v, ok := parsed.Flags["sort"]; ok {
		v = strings.ToLower(v)
		if v != "asc" && v != "desc" {
			return fmt.Errorf("invalid --sort: must be 'asc' or 'desc'")
		}
		sortOrder = v
	}

	// --filter <string>
	if v, ok := parsed.Flags["filter"]; ok {
		filterQ = v
	}

	// If sorting/filtering, fetch more then trim locally.
	fetchLimit := int32(limit)
	if filterQ != "" || (parsed.Flags["sort"] != "") {
		const maxFetch = 100
		fetchLimit = int32(maxFetch)
	}

	posts, err := s.DB.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  fetchLimit,
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	// Filter (case-insensitive: checks title, description, feed name, url)
	if filterQ != "" {
		q := strings.ToLower(filterQ)
		filtered := make([]database.GetPostsForUserRow, 0, len(posts))
		for _, p := range posts {
			title := strings.ToLower(p.Title)
			desc := strings.ToLower(p.Description.String)
			feed := strings.ToLower(p.FeedName)
			url := strings.ToLower(p.Url)
			if strings.Contains(title, q) || strings.Contains(desc, q) ||
				strings.Contains(feed, q) || strings.Contains(url, q) {
				filtered = append(filtered, p)
			}
		}
		posts = filtered
	}

	// Sort by PublishedAt (asc = oldest first, desc = newest first)
	sort.Slice(posts, func(i, j int) bool {
		ti := posts[i].PublishedAt.Time
		tj := posts[j].PublishedAt.Time
		if sortOrder == "asc" {
			return ti.Before(tj)
		}
		return tj.Before(ti)
	})

	// Trim to limit
	if limit > 0 && len(posts) > limit {
		posts = posts[:limit]
	}

	fmt.Println("Flags:", parsed.Flags)
	fmt.Println("Positionals:", parsed.Positionals)
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
