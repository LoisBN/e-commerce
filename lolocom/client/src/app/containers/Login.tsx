import React from "react";
import { Field, reduxForm } from "redux-form";

import { renderField, renderFieldDate } from "../components/Login";

const validate = (values: any) => {
	const Day = new Date();
	const maxDay = [
		Day.getFullYear() - 18,
		Day.getMonth() + 1,
		Day.getDate() + 1
	];
	const errors: { email?: string; password?: string; birth?: string } = {};
	if (!values.email) {
		errors.email = "Required";
	} else if (
		!/^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i.test(values.email)
	) {
		errors.email = "Invalid email address";
	}
	if (!values.password) {
		errors.password = "Required";
	} else if (values.password.length < 8) {
		errors.password = "Must be 8 characters or more";
	} else if (values.password.length > 28) {
		errors.password = "Must be 28 characters or less";
	}
	if (!values.birth) {
		errors.birth = "Required";
	} else if (
		Number(values.birth.split("-")[0]) > maxDay[0] ||
		(Number(values.birth.split("-")[0]) === maxDay[0] &&
			Number(values.birth.split("-")[1]) > maxDay[1]) ||
		(Number(values.birth.split("-")[0]) === maxDay[0] &&
			Number(values.birth.split("-")[1]) === maxDay[1] &&
			Number(values.birth.split("-")[2]) > maxDay[2])
	) {
		errors.birth = "You must be over 18";
	}
	return errors;
};

const warn = (values: any) => {
	const warnings: any = {};
	/*
	if (values.birth < 18) {
		warnings.birth = "Hmm, you seem a bit young...";
	}
	*/
	return warnings;
};

const Login = ({ ...props }) => {
	const { handleSubmit } = props;
	const Day = new Date();
	const maxDay = [
		Day.getFullYear() - 18,
		Day.toLocaleDateString("en-CA").slice(5, 10)
	].join("-");
	return (
		<div>
			<img alt={""} src={""} />
			<form onSubmit={val => handleSubmit(val)}>
				<button>Connect With Google</button>
				<button>Connect With Facebook</button>
				<div>
					<Field
						label="Email"
						name="email"
						component={renderField}
						type="email"
					/>
				</div>
				<div>
					<Field
						label="Password"
						name="password"
						component={renderField}
						type="password"
					/>
				</div>
				<div>
					<Field
						label="naissance"
						name="birth"
						component={renderFieldDate}
						type="date"
						max={maxDay}
						defaultValue={maxDay}
					/>
				</div>
				<button type="submit">Submit</button>
			</form>
		</div>
	);
};

export const LoginPage = reduxForm({
	form: "LOGIN",
	validate,
	warn
})(Login);

export const LogIcon = () => {
	return true ? (
		<i className="far fa-user"></i>
	) : (
		<i className="fas fa-user"></i>
	);
};
