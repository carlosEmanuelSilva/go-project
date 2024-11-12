import React, { Component } from "react";
import "./Input.scss";

class Input extends Component {
	render() {
		return (
			<div className="input">
				<input onKeyDown={this.props.send} />
			</div>
		);
	}
}

export default Input;
