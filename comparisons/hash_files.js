const fs = require("fs");
const crypto = require("crypto");

async function fileHash(filename, algorithm = "sha256") {
  return new Promise((resolve, reject) => {
    let shasum = crypto.createHash(algorithm);
    try {
      let s = fs.ReadStream(filename)
      s.on("data", function (data) {
        shasum.update(data)
      })
      // making digest
      s.on("end", function () {
        const hash = shasum.digest("hex")
        return resolve(hash);
      })
    } catch (error) {
      return reject("calc fail");
    }
  });
}

async function hashNTimes(n) {
  function hashi(index) {
    return fileHash("./cat.jpg");
  }

  var ps = [];
  for (var i = 0; i < n; i++) {
    ps.push(hashi(i));
  }

  return Promise.all(ps);
}

const NS_PER_SEC = 1e9;
const time = process.hrtime();

hashNTimes(1000)
  .then(() => {
    const diff = process.hrtime(time);
    console.log(`Execution time = ${diff[0] * NS_PER_SEC + diff[1]} nanoseconds`);
  });
