# Sorting stability

This sample application highlights a situation when sorting stability is needed.

This data will be sorted twice:

1. Sort by `Name`
1. Sort by `State`

If the end result is the dataset being sorted by state and name, the first sort of `Name` does *not* require stability. This demo highlights the output though of the second sort:

* If you sort with an **unstable** algorithm then you can lose the inner sorting by `Name`
* If you sort with a **stabe** algorithmn, you preserve the inner sorting by `Name`, so the output dataset is sorted by state and name correctly

## Output

```text
--- Unsorted ----
Name: Thomas, State: New Hampshire
Name: Annie, State: New Hampshire
Name: Chris, State: Maine
Name: Jackie, State: Vermont
Name: Marie, State: New Hampshire
Name: Daniel, State: Maine
Name: Vivian, State: Maine
Name: Michelle, State: Maine

--- Sorted by Name (heap sort: unstable) ----
Name: Annie, State: New Hampshire
Name: Chris, State: Maine
Name: Daniel, State: Maine
Name: Jackie, State: Vermont
Name: Marie, State: New Hampshire
Name: Michelle, State: Maine
Name: Thomas, State: New Hampshire
Name: Vivian, State: Maine

--- Sorted by State (heap sort: unstable) ----
Name: Daniel, State: Maine
Name: Michelle, State: Maine
Name: Chris, State: Maine
Name: Vivian, State: Maine
Name: Thomas, State: New Hampshire
Name: Marie, State: New Hampshire
Name: Annie, State: New Hampshire
Name: Jackie, State: Vermont

************************************
Now trying with a stable sort on state
************************************

--- Unsorted ----
Name: Thomas, State: New Hampshire
Name: Annie, State: New Hampshire
Name: Chris, State: Maine
Name: Jackie, State: Vermont
Name: Marie, State: New Hampshire
Name: Daniel, State: Maine
Name: Vivian, State: Maine
Name: Michelle, State: Maine

--- Sorted by Name (heap sort: unstable) ----
Name: Annie, State: New Hampshire
Name: Chris, State: Maine
Name: Daniel, State: Maine
Name: Jackie, State: Vermont
Name: Marie, State: New Hampshire
Name: Michelle, State: Maine
Name: Thomas, State: New Hampshire
Name: Vivian, State: Maine

--- Sorted by State (merge sort: stable) ----
Name: Chris, State: Maine
Name: Daniel, State: Maine
Name: Michelle, State: Maine
Name: Vivian, State: Maine
Name: Annie, State: New Hampshire
Name: Marie, State: New Hampshire
Name: Thomas, State: New Hampshire
Name: Jackie, State: Vermont
```
