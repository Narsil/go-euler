#!/bin/sh
cd profiling && make install && cd ..
cd primes && make install && cd ..
cd arithmetics && make install && cd ..
cd bigextension && make install && cd ..
cd problems && make install && cd ..
cd cmd && make install && ./euler
