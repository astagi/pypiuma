import os
import time

from pypiuma import optimize, PiumaError

start = time.time()

for image_file_name in os.listdir("images"):
	try:
		path = optimize(os.path.join("images", image_file_name), 100, 50)
	except PiumaError as err:
		print (err)

end = time.time()
print(end - start)
