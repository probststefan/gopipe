<?php

echo "Hallo PHP!";

file_put_contents('newfile.txt', "hallo".PHP_EOL , FILE_APPEND | LOCK_EX);
