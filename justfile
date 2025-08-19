# Exercism Repository Management

# Show summary of PR activity
summary:
    @just pr-calendar

# Show PR closure activity in a calendar format
pr-calendar:
    #!/usr/bin/env bash
    echo "📊 PR Closure Activity Calendar" | gum style --foreground 212 --bold
    echo
    
    # Get PR closure data
    gh pr list --state closed --limit 200 --json closedAt,number,title | \
    jq -r '.[] | select(.closedAt != null) | .closedAt[:10]' | \
    sort | uniq -c | \
    awk '{print $2 " " $1}' > /tmp/pr_dates.txt
    
    # Get current date info
    current_year=$(date +%Y)
    current_month=$(date +%m)
    today=$(date +%Y-%m-%d)
    
    # Generate calendar for current month
    echo "$(date +"%B %Y")" | gum style --foreground 39 --bold --align center
    echo
    
    # Get first day of month and number of days
    first_day=$(date -d "$current_year-$current_month-01" +%u)
    if [ "$first_day" -eq 7 ]; then first_day=0; fi
    days_in_month=$(date -d "$current_year-$current_month-01 +1 month -1 day" +%d)
    
    # Build calendar data for gum table
    calendar_data="Sun,Mon,Tue,Wed,Thu,Fri,Sat"
    
    # Initialize week with empty cells
    week=("" "" "" "" "" "" "")
    day=1
    
    # Fill in the calendar weeks
    while [ $day -le $days_in_month ] || [ ${#week[@]} -gt 0 ]; do
        # Start a new week
        for col in {0..6}; do
            if [ $day -le $days_in_month ] && [ $((($day + $first_day - 1) % 7)) -eq $col ]; then
                # Build date string carefully
                if [ $day -lt 10 ]; then
                    date_str="$current_year-$current_month-0$day"
                else
                    date_str="$current_year-$current_month-$day"
                fi
                
                count=$(grep "^$date_str " /tmp/pr_dates.txt 2>/dev/null | awk '{print $2}' || echo "0")
                
                # Ensure count is a valid number
                case "$count" in
                    ''|*[!0-9]*) count="0" ;;
                esac
                
                # Format the day with symbols - underline today
                if [ "$date_str" = "$today" ]; then
                    if [ "$count" -gt 10 ]; then
                        week[$col]="🔥$(printf '\033[4m%s\033[0m' $day)"
                    elif [ "$count" -gt 5 ]; then
                        week[$col]="🚀$(printf '\033[4m%s\033[0m' $day)"
                    elif [ "$count" -gt 0 ]; then
                        week[$col]="✅$(printf '\033[4m%s\033[0m' $day)"
                    else
                        week[$col]="$(printf '\033[4m%s\033[0m' $day)"
                    fi
                elif [ "$count" -gt 10 ]; then
                    week[$col]="🔥$day"
                elif [ "$count" -gt 5 ]; then
                    week[$col]="🚀$day"
                elif [ "$count" -gt 0 ]; then
                    week[$col]="✅$day"
                else
                    week[$col]="$day"
                fi
                
                day=$((day + 1))
            elif [ $day -le $days_in_month ]; then
                week[$col]=""
            else
                week[$col]=""
            fi
        done
        
        # Add this week to calendar data
        calendar_data="$calendar_data\n${week[0]},${week[1]},${week[2]},${week[3]},${week[4]},${week[5]},${week[6]}"
        
        # Check if we're done
        if [ $day -gt $days_in_month ]; then
            break
        fi
    done
    
    # Display the calendar using gum table (print mode to avoid clearing screen)
    echo -e "$calendar_data" | gum table --widths 6,6,6,6,6,6,6 --print
    
    echo
    echo "Legend:" | gum style --foreground 39 --bold
    echo "🔥 10+ PRs   🚀 6-10 PRs   ✅ 1-5 PRs" | gum style --foreground 242
    
    # Show detailed daily activity for this month
    echo
    echo "Daily PR Activity:" | gum style --foreground 39 --bold
    
    # Build detailed activity table
    activity_data="Date,Day,PRs Closed"
    
    # Get all dates with PR activity this month and format them
    grep "^$current_year-$current_month" /tmp/pr_dates.txt 2>/dev/null | while read date count; do
        if [ -n "$date" ] && [ -n "$count" ]; then
            day_name=$(date -d "$date" +%a)
            formatted_date=$(date -d "$date" +"%b %d")
            
            # Add visual indicator
            if [ "$count" -gt 10 ]; then
                indicator="🔥"
            elif [ "$count" -gt 5 ]; then
                indicator="🚀"
            elif [ "$count" -gt 0 ]; then
                indicator="✅"
            fi
            
            # Check if it's today
            if [ "$date" = "$today" ]; then
                formatted_date="$(printf '\033[4m%s\033[0m' "$formatted_date")"
            fi
            
            echo "$formatted_date,$day_name,$indicator $count"
        fi
    done | sort -t, -k1M -k1d > /tmp/activity_table.txt
    
    # Only show table if there's data
    if [ -s /tmp/activity_table.txt ]; then
        (echo "$activity_data"; cat /tmp/activity_table.txt) | gum table --print
    else
        echo "No PR activity this month" | gum style --foreground 242
    fi
    
    # Show this month's stats
    echo
    echo "This Month's Summary:" | gum style --foreground 39 --bold
    total_this_month=$(grep "^$current_year-$current_month" /tmp/pr_dates.txt | awk '{sum += $2} END {print sum+0}')
    active_days=$(grep "^$current_year-$current_month" /tmp/pr_dates.txt | wc -l)
    echo "Total PRs closed: $total_this_month" | gum style --foreground 46
    echo "Active days: $active_days" | gum style --foreground 46
    if [ "$active_days" -gt 0 ]; then
        avg_prs=$(echo "scale=1; $total_this_month / $active_days" | bc -l 2>/dev/null || echo "0")
        echo "Average PRs per active day: $avg_prs" | gum style --foreground 46
    fi
    
    rm -f /tmp/pr_dates.txt /tmp/activity_table.txt

# Show detailed PR stats
stats:
    @echo "📈 Detailed PR Statistics" | gum style --foreground 212 --bold
    @echo
    @echo "Total PRs (all time):" | gum style --foreground 39
    @gh pr list --state all --limit 1000 | wc -l
    @echo
    @echo "Open PRs:" | gum style --foreground 46
    @gh pr list --state open | wc -l
    @echo
    @echo "Closed PRs:" | gum style --foreground 196
    @gh pr list --state closed --limit 1000 | wc -l
    @echo
    @echo "Most recent PRs:" | gum style --foreground 39
    @gh pr list --state all --limit 5

# Show today's PR activity
today:
    #!/usr/bin/env bash
    today=$(date +%Y-%m-%d)
    echo "📅 Today's PR Activity ($today)" | gum style --foreground 212 --bold
    echo
    
    closed_today=$(gh pr list --state closed --limit 100 --json closedAt,number,title | \
        jq -r ".[] | select(.closedAt | startswith(\"$today\")) | \"#\(.number) - \(.title)\"")
    
    if [ -z "$closed_today" ]; then
        echo "No PRs closed today" | gum style --foreground 242
    else
        echo "PRs closed today:" | gum style --foreground 46
        echo "$closed_today" | gum style --foreground 39
    fi

# Merge all open PRs (use with caution!)
merge-all:
    @echo "⚠️  This will merge ALL open PRs!" | gum style --foreground 196 --bold
    @gum confirm "Are you sure you want to continue?" && \
        gh pr list --json number | jq -r '.[].number' | xargs -I {} gh pr merge {} --squash --delete-branch

# Show help
help:
    @just --list