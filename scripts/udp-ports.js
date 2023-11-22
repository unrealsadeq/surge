var portRanges = [
  { start: 5222, end: 5222 },
  { start: 7085, end: 8200 },
  { start: 8700, end: 8700 },
  { start: 9030, end: 9030 },
  { start: 17000, end: 19301 },
  { start: 19303, end: 20100 }
];

for (var i = 0; i < portRanges.length; i++) {
  var range = portRanges[i];
  if (port >= range.start && port <= range.end) {
    portMatched = true;
    break;
  }
}

$done({ matched: portMatched });