import React from "react";
import { NewsLettersPres } from "../../../components/Communs/Footer/NewsLetters";

import { reduxForm, Field } from "redux-form";

const NewsLettersContainer = () => {
	const signToNewsletter = (mail: any) => {
		console.log(mail);
	};
	return (
		<NewsLettersPres>
			<form onSubmit={mail => signToNewsletter(mail)}>
				<Field
					type="email"
					name="email"
					label="email"
					component={NewsMail}
				/>
				<button type="submit">Subscribe</button>
			</form>
		</NewsLettersPres>
	);
};

export const NewsLetters = reduxForm({ form: "signToNewsletter" })(
	NewsLettersContainer
);

const NewsMail = ({
	input,
	label,
	type,
	meta: { touched, error, warning }
}: {
	input: any;
	label: string;
	type: string;
	meta: { touched: any; error: any; warning: any };
}) => (
	<div>
		<label>{label}</label>
		<div>
			<input {...input} placeholder={label} type={type} />
			{touched &&
				((error && <span>{error}</span>) ||
					(warning && <span>{warning}</span>))}
		</div>
	</div>
);
