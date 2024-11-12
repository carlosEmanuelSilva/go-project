import React { Component } from "react";
import "./Inpust.scss";

class Input extends Component {
	render() {
		return (
			<div className="input">
				<input onKeyDown={this.prop.send} />
			</div>
		);
	}
}

export default Input;
