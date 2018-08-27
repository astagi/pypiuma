import os
import time

from pypiuma import optimize, optimize_dir, optimize_list, PiumaError

images_to_optimize = [ os.path.join("images", item) for item in os.listdir("images") for repetitions in range(1) ]

start = time.time()

for image_file_name in images_to_optimize:
	try:
		path = optimize(image_file_name, 100, 50)
	except PiumaError as err:
		print (err)

end = time.time()
print(end - start)

start = time.time()

try:
	path = optimize_list(images_to_optimize, 100, 50)
except PiumaError as err:
	print (err)

end = time.time()

print(end - start)

start = time.time()

try:
	path = optimize_dir("./images", 100, 50)
except PiumaError as err:
	print (err)

end = time.time()


print(end - start)
