/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 09-10-2018
 * |
 * | File Name:     avidnet.js
 * +===============================================
 */
var cfg = {
  _id: "avidnet",
  members: [
    {
      _id: 0,
      host: "192.168.1.88:27017"
    },
    {
      _id: 1,
      host: "192.168.1.81:27017"
    },
    {
      _id: 2,
      arbiterOnly: true,
      host: "192.168.1.88:27018"
    }
  ]
};
rs.initiate(cfg, { force: true });
rs.config();
