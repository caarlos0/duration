# duration

Copy of stdlib's `time.Duration`, but `ParseDuration` accepts other units as well:

- `d`: days (`24 * time.Hour`)
- `w`: weeks (`7 * Day`)
- `mo`: months (`30 * Day`)
- `y`: years (`365 * Day`)

This is specially useful if you want to accept duration as flags in your program, and those duration are expected to be high.
Asking the user to type `7860h` instead of `1y` might be something you don't want to.
If so, you can accept the flag as string and parse with this package instead.

Hopefully, someday, something like this gets merged into the stdlib.
