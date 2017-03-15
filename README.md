# gobotics
Golang robotics apis and libraries<br>

Gobotics is middleware that provides a policy interfaces to robotic devices.<br><br>

Gobotics provides DB services, secure communication layers, and policy administration.<br> 
Providing clean interfaces with command and control of robotic communities and swarms through<br>
easy to implement APIs.<br><br>

Gobotics can be used as director level software services to control bots from shop floor, through<br>
the supply chain, all the way to customer delivery and billing.<br><br>

You need GCC-7.0 and go1.7 to compile the current implementation this include gobotics and gobot-io<br>
packages. This has been tested on Ubuntu 14.04.<br><br>

Make a clean directory:
mkdir <somenewgccdir> <br>
cd <somenewgccdir>
svn checkout svn://gcc.gnu.org/svn/gcc/trunk . <br><br>

download build and install flex-2.6.0<br>
https://sourceforge.net/projects/flex/<br>
./configure <br>
make; make install <br>

apt-get install bison<br>

make sure the flex-2.6.0 is in your path<br>

mkdir <someothergccconfigdir_on_different_level_from_new_gcc> <br>
mkdir <someoutputdir> <br>

cd someothergccconfigdir_on_different_level_from_new_gcc<br>
../somenewgccdir/configure --prefix=someothergccconfigdir/someoutputdir --enable-languages=c,c++,go --disable-multilib<br>
Put the new someoutputdir/bin in your path.<br>

TBD: Need to fix static line of libgo with proper export LD_LIBRARY_PATH 