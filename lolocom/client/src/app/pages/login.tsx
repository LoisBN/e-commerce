import React from "react";
import { LoginPage } from "../containers/Login";

export const Login = () => {
	const handleSubmit = (values: any) => {
		console.log(values);
	};
	return <LoginPage onSubmit={handleSubmit} />;
};
