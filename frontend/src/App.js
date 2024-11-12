import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import Input from "./components/Input/Input";

class App extends Component {
  constructor(props) {
    super(props);
	this.state = {
		chatHistory: []
	}
  }

	componentDidMount() {
		connect((msg) => {
			console.log("Nova mensagem")
			this.setState(prevState => ({
				chatHistory: [...this.state.chatHistory, msg]
			}))

			console.log(this.state);
		});
	}

  send(event) {
	  if(event.keyCode === 13) {
		sendMsg(event.target.value);
		event.target.value = "";
	  }
  }

  render() {
	return (
		<div className="App">
			<Header />
			<ChatHistory chatHistory = {this.state.chatHistory} />
			<Input send={this.send} />
		</div>
	);
  }
}

export default App;
