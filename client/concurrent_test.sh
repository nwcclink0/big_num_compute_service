#!/bin/bash

for i in {1..10000} ; do
  ./client -o create &
  ./client -o update &
  ./client -o delete &
done

./client -o create_me&
./client -o create_cat&

for i in {1..10000} ; do
./client -o compute_me_cat_add &
./client -o compute_me_cat_sub &
./client -o compute_me_cat_mul &
./client -o compute_me_cat_div &
./client -o compute_me_add_random &
./client -o compute_me_sub_random &
./client -o compute_me_mul_random &
./client -o compute_me_div_random &
done

#./client -o delete_me
#./client -o delete_cat
