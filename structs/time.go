package structs

import (
	"fmt"
	"time"
)

func TimeAgo(t time.Time) string {
	diff := time.Since(t)

	// Calculate the number of years, months, days, hours, and minutes
	years := int(diff.Hours()) / (24 * 365)
	months := int(diff.Hours()) / (24 * 30) // Approximating months as 30 days
	days := int(diff.Hours()) / 24          // Total days
	hours := int(diff.Hours()) % 24         // Remainder hours after calculating days
	minutes := int(diff.Minutes()) % 60     // Remainder minutes after calculating hours

	// If the difference is over a year
	if years > 0 {
		if years == 1 {
			return "1 year ago"
		}
		return fmt.Sprintf("%d years ago", years)
	}

	// If the difference is over a month
	if months > 0 {
		if months == 1 {
			return "1 month ago"
		}
		return fmt.Sprintf("%d months ago", months)
	}

	// If the difference is over a day
	if days > 0 {
		if days == 1 {
			return "1 day ago"
		}
		return fmt.Sprintf("%d days ago", days)
	}

	// If the difference is less than a day, check hours
	if hours > 0 {
		if hours == 1 {
			return "1 hour ago"
		}
		return fmt.Sprintf("%d hours ago", hours)
	}

	// If the difference is less than an hour, check minutes
	if minutes > 0 {
		if minutes == 1 {
			return "1 minute ago"
		}
		return fmt.Sprintf("%d minutes ago", minutes)
	}

	// If it's less than a minute, return "Just now"
	return "Just now"
}
