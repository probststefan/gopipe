const fs = require('fs');
fs.writeFile("newfile.txt", "Hey there from js!\n", function(err) {
    if(err) {
        return console.log(err);
    }

    console.log("The file was saved!");
});
