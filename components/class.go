package components

import "strings"

const (
	Card            = "rounded-lg border border-black dark:border-white bg-card text-card-foreground shadow-sm"
	CardHeader      = "flex flex-col space-y-1.5 p-6"
	CardTitle       = "text-2xl font-semibold leading-none tracking-tight"
	CardDescription = "text-sm text-muted-foreground"
	CardContent     = "p-6 pt-0"
	CardFooter      = "flex items-center p-6 pt-0"

	Input = "flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-base ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm"
	Label = "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"

	NavigationMenu     = "relative z-10 flex max-w-max flex-1 items-center justify-center my-1"
	NavigationMenuList = "group flex flex-1 list-none items-center justify-center space-x-1"

	DropDownMenu        = "flex cursor-default gap-2 select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none focus:bg-accent data-[state=open]:bg-accent [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0"
	DropDownMenuContent = "z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2"
	DropDownLabel       = "px-2 py-1.5 text-sm font-semibold"
	DropDownMenuItem    = "relative flex cursor-default select-none items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0"
)

func ButtonVariant(variant string) string {
	if variant == "" {
		variant = "default"
	}

	switch variant {
	case "destructive":
		return "bg-destructive text-destructive-foreground hover:bg-destructive/90"
	case "outline":
		return "border border-input bg-background hover:bg-accent hover:text-accent-foreground"
	case "secondary":
		return "bg-secondary text-secondary-foreground hover:bg-secondary/80"
	case "ghost":
		return "hover:bg-accent hover:text-accent-foreground"
	case "link":
		return "text-primary underline-offset-4 hover:underline"
	case "default":
		return "bg-primary text-primary-foreground hover:bg-primary/90"
	default:
		return "bg-primary text-primary-foreground hover:bg-primary/90"
	}
}

func ButtonSize(size string) string {
	if size == "" {
		size = "default"
	}

	switch size {
	case "icon":
		return "h-6 w-8"
	case "lg":
		return "h-11 rounded-md px-8"
	case "sm":
		return "h-9 rounded-md px-3"
	case "default":
		return "h-10 px-4 py-2"
	default:
		return "h-10 px-4 py-2"
	}
}

// classesToString converts a slice of classes into a cleaned, space-separated string
func classesToString(classList []string) string {
	if len(classList) == 0 {
		return ""
	}

	// Use a map to remove duplicates and handle redundancy
	classMap := make(map[string]bool)

	// Optional: Define redundancy rules (expand as needed)
	redundancies := map[string][]string{
		"flex": {"flex-col", "flex-row"}, // flex is redundant if direction is specified
	}

	for _, cls := range classList {
		// Skip if this class is redundant based on existing classes
		if redundantWith, exists := redundancies[cls]; exists {
			for _, redundant := range redundantWith {
				if classMap[redundant] {
					continue // Skip adding "flex" if "flex-col" or "flex-row" is present
				}
			}
		}
		classMap[cls] = true
	}

	// Convert map keys back to a slice
	var finalClasses []string
	for cls := range classMap {
		finalClasses = append(finalClasses, cls)
	}

	// Join into a single string with spaces
	return strings.Join(finalClasses, " ")
}

// cleanClasses removes redundant or conflicting CSS classes from a space-separated string
func CleanClasses(classList []string) string {
	classString := classesToString(classList)

	// Split into individual classes
	classes := strings.Fields(classString)
	if len(classes) <= 1 {
		return classString // Nothing to clean
	}

	// Use a map to track seen classes, but preserve order
	seen := make(map[string]bool)
	var uniqueClasses []string

	// Add classes in order, skipping duplicates
	for _, cls := range classes {
		if !seen[cls] {
			seen[cls] = true
			uniqueClasses = append(uniqueClasses, cls)
		}
	}

	// Join back into a string
	return strings.Join(uniqueClasses, " ")
}
