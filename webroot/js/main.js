define(
	"main",
	[
		"MessageList"
	],
	function(MessageList) {
		var ws = new WebSocket("ws://localhost:80/entry");
		var list = new MessageList(ws);
		ko.applyBindings(list);
	}
);
