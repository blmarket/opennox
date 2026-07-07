See ./t1 for context.

Check current coverage, and find one function which is safe to convert from C to Go.

Make the migration.

`git commit` with proper title and description

- Try to clarify type sizes - Use C.uint32_t instead of C.uint
- Pick a function which is already covered well with existing test cases. Ignore
  if coverage is insufficient.

