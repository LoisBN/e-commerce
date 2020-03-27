import React from "react";
import { Field, reduxForm } from "redux-form";

import { renderField, renderFieldDate, LoginPres } from "../components/Login";
import { AuthModeSwitcher } from "../components/Switchers";

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
	const [isKnown, setIsKnown] = React.useState<boolean>(false);
	const Day = new Date();
	const maxDay = [
		Day.getFullYear() - 18,
		Day.toLocaleDateString("en-CA").slice(5, 10)
	].join("-");
	const switchIsKnown = () => {
		setIsKnown(prev => !prev);
	};
	return (
		<LoginPres>
			<AuthModeSwitcher
				setter={switchIsKnown}
				current={isKnown}
				texts={["Sign Up", "Sign In"]}
			/>
			<button>Connect With Google</button>
			<button>Connect With Facebook</button>
			<form onSubmit={val => handleSubmit(val)}>
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
				{!isKnown && (
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
				)}
				<button type="submit">Submit</button>
			</form>
		</LoginPres>
	);
};

export const LoginPage = reduxForm({
	form: "LOGIN",
	validate,
	warn
})(Login);

export const LogIcon = () => {
	return (
		<button>
			{true ? (
				<i className="far fa-user"></i>
			) : (
				<i className="fas fa-user"></i>
			)}
		</button>
	);
};
