require "fileutils"

desc "compile all files"
task :all do

  print "compiling"
  Dir["./*.go"].each { |f| 
    print(".")
    system("go build #{f}")

    sf = f.split("")
    3.times { sf.pop }

    FileUtils.move("#{sf.join("")}","./bins/",force: true)
    print(".")
  }
  print(".  OK\n")

end
