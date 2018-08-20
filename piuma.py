import os
import time

from pypiuma import optimize, PiumaError

images_to_optimize = [ item for item in os.listdir("images") for repetitions in range(400) ]

start = time.time()

for image_file_name in images_to_optimize:
	try:
		path = optimize(os.path.join("images", image_file_name), 100, 50)
	except PiumaError as err:
		print (err)

end = time.time()
print(end - start)
