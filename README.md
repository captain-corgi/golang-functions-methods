# Assessment

## https://www.coursera.org/learn/golang-functions-methods/peer/qKrnv/module-2-activity/submit

Let us assume the following formula for
displacement *s* as a function of time *t*, acceleration *a*, initial velocity *v*~*o*~,
and initial displacement *s*~*o*~.

*s *=½* a t*^*2*^+* v*~*o*~*t *+* s*~*o*~

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity v~o~, and initial
displacement s~o~. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: *a* = 10, *v*~*o*~ = 2, *s*~*o*~ = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
