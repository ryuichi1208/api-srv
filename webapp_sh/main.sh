#!/bin/bash

nc -kl 3000 -w 1 0<backpipe | awk '/HTTP/{system("cat " substr($2, 2))}' 1>backpipe
