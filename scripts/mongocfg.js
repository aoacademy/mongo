/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 13-04-2018
 * |
 * | File Name:     scripts/mongocfg.js
 * +===============================================
 */
rs.initiate(
  {
    _id: "cfgrepl",
    configsvr: true,
    members: [
      { _id : 0, host : "mongocfg1:27017" },
    ]
  }
)
