from z3 import *
import re

# Read input-user file
f = open("input-user.txt", "r")
lines = f.readlines()

# Compile regexp and find digits
regexp_digits = re.compile(r'\d+')
digits = regexp_digits.findall(",".join(lines))

# Process digits in chunks of 6 as in ax, ay, bx, by, rx, and ry
total_points = 0
for i in range(0, len(digits), 6):
    ax = digits[i]
    ay = digits[i+1]
    bx = digits[i+2]
    by = digits[i+3]
    rx = int(digits[i + 4])
    ry = int(digits[i + 5])

    # For part 2, we add a big factor ;)
    rx += 10000000000000
    ry += 10000000000000

    # Create integer variables for the number of times to press each button
    a = Int('a')
    b = Int('b')

    # Create the solver
    s = Solver()

    # Define the constraints
    s.add(ax * a + bx * b == rx)  # X-coordinate constraint
    s.add(ay * a + by * b == ry)  # Y-coordinate constraint

    # Add the condition for minimization
    cost = 3 * a + b

    # Fetch min cost and loop to see if we can find anything with lower cost
    min_cost = None
    while s.check() == sat:
        model = s.model()
        current_cost = model.evaluate(cost).as_long()
        if min_cost is None or current_cost < min_cost:
            min_cost = current_cost

        # Add a constraint to find a better solution than the current model
        s.add(cost < current_cost)

    if min_cost != None:
        total_points += min_cost

print("Total points:", total_points)
