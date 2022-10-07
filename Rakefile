require "fileutils"

desc "compile all files"
task :all do

  print "compiling"
  Dir["./src/*.go"]::each { |f| 
    print(".")
    system("go build #{f}")

    sf = f.gsub("./src/","")
      .gsub(".go","")
    p sf  


    FileUtils.move("#{sf}","./bins",force: true)
    print(".")
  }
  print(".  OK\n")

end
