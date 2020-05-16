import * as mql from "./mql_";
import axios from "axios";

let c = new mql.Client({
  host: "https://metamate.one",
  client: axios.create(),
});

const run = async () => {
  let rsp = await c.GetPosts({
    serviceFilter: {
      id: {
        value: {
          is: "hackernews",
        }
      }
    },
    mode: {
      kind: mql.GetModeKind.Search,
      search: {
        term: "book recommendations",
      },
    },
  });

  console.log(rsp);
};

run();