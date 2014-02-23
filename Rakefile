GOSRC = Dir.pwd + "/go/src"
GOPATH = Dir.pwd + "/go"
GOCMD = "/usr/local/go/bin/go"

task :buildmodule  do
	
	Dir.chdir("./cpp") do
		sh "swig -go -c++ -intgosize 64 helloworld.swigcxx"
		sh "g++ -c -fPIC helloworld.cxx"
		sh "g++ -c -fPIC helloworld_wrap.cxx"
		sh "g++ -undefined suppress -flat_namespace -shared helloworld.o helloworld_wrap.o -o helloworld.so"
		sh "cp helloworld.go #{GOPATH}/src/helloworld/" 
		sh "cp helloworld_gc.c #{GOPATH}/src/helloworld/"
		sh "cp helloworld.so #{GOPATH}/src/"
	end

	Dir.chdir("#{GOPATH}/src/helloworld") do
		sh "GOPATH=#{GOPATH} #{GOCMD} tool 6c -I /usr/local/go/pkg/darwin_amd64 -D _64BIT helloworld_gc.c"
		sh "GOPATH=#{GOPATH} #{GOCMD} tool 6g helloworld.go"
		sh "GOPATH=#{GOPATH} #{GOCMD} tool pack grc helloworld.a helloworld_gc.6 helloworld.6"
		sh "GOPATH=#{GOPATH} #{GOCMD} install"
	end

end

task :build do 

	Dir.chdir("#{GOSRC}") do
		puts `GOPATH=#{GOPATH} #{GOCMD} build -o hellotester`
	end	

end