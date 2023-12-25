from z3 import *

x = Real('x')
y = Real('y')
z = Real('z')
a = Real('a')
b = Real('b')
c = Real('c')
t = Real('t')
r = Real('r')
s = Real('s')
solver = Solver()
solver.add(0 == x + a * t - 246721424318191 - 46 * t)
solver.add(0 == y + b * t - 306735195971895 + 42 * t)
solver.add(0 == z + c * t - 195640804079938 - 141 * t)
solver.add(0 == x + a * r - 286716952521568 - 121 * r)
solver.add(0 == y + b * r - 348951612232772 - 421 * r)
solver.add(0 == z + c * r - 274203424013154 + 683 * r)
solver.add(0 == x + a * s - 231402843137765 - 30 * s)
solver.add(0 == y + b * s - 83297412652001 - 154 * s)
solver.add(0 == z + c * s - 273065723902291 - 66 * s)
solver.add(t > 0)
solver.add(r > 0)
solver.add(s > 0)
if solver.check() == sat:
    m = solver.model()
    print(int(str(m.evaluate(x)))+int(str(m.evaluate(y)))+int(str(m.evaluate(z))))
else:
    print("wrong")