from purepiuma import pureppiuma
from pypiuma import optimize_dir


def main():
    pureppiuma()
    optimize_dir("./images", 200)


if __name__== "__main__":
  main()
