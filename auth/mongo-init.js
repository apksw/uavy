db.createUser(
  {
    user: "auth",
    pwd: "auth",
      roles: [
        {
          role: "readWrite",
          db: "auth"
        }
      ]
  }
);
