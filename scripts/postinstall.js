var mkdirp = require('mkdirp');
var path = require('path');
var ncp = require('ncp');

// Paths
var ns = 'ntoolkit';
var src = path.join(__dirname, '..', 'src', ns);
var dir = path.join(__dirname, '..', '..', '..', 'src', 'vendor', ns);
var ismodule = __dirname.split(path.sep).filter(function(i) { return i == 'node_modules'; }).length > 0;

// Create folder if missing
if (ismodule) {
  mkdirp(dir, function (err) {
    if (err) {
      console.error(err)
      process.exit(1);
    }

    // Copy files
    ncp(src, dir, function (err) {
      if (err) {
        console.error(err);
        process.exit(1);
      }
    });
  });
}
