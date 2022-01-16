import Mongoose from "mongoose";

let database: Mongoose.Connection;
export const connect = (): Mongoose.Connection => {
  const uri = process.env.MONGO_CONN_STRING!;

  console.log("trying to connect");
  if (database) {
    return database;
  }
  Mongoose.connect(uri);
  database = Mongoose.connection;
  database.once("open", async () => {
    console.log("Connected to database");
  });
  database.on("error", () => {
    console.log("Error connecting to database");
  });

  return database;
};

export const disconnect = () => {
  if (!database) {
    return;
  }
  Mongoose.disconnect();
};
