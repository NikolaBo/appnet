internal{}

fn init() {}

fn req(rpc_req) {
  rpc_req.set("language", "english");
  rpc_req.set("location", "seattle");
}

fn resp(rpc_resp) {
	send(rpc_resp, APP);
}