# csaw demo recording
# Re-record: vhs docs/assets/demo.tape

Output docs/assets/demo.gif
Set Theme "Catppuccin Mocha"
Set FontSize 16
Set Width 900
Set Height 520
Set Padding 20
Set TypingSpeed 40ms

# Show sources
Type "csaw source list"
Enter
Sleep 1200ms

# Mount with a profile
Enter
Type "mkdir -p .claude && csaw mount --profile dotagent/go --force"
Enter
Sleep 1800ms

# Inspect the result
Enter
Type "csaw inspect"
Enter
Sleep 3000ms

# Health check
Enter
Type "csaw check"
Enter
Sleep 1200ms

# Status overview
Enter
Type "csaw status"
Enter
Sleep 1500ms

# Clean unmount
Enter
Type "csaw unmount"
Enter
Sleep 1200ms

Sleep 800ms
