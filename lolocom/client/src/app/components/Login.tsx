import React from "react";

export const renderField = ({
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

export const renderFieldDate = ({
	input,
	label,
	max,
	meta: { touched, error, warning }
}: {
	input: any;
	label: string;
	max: string;
	meta: { touched: any; error: any; warning: any };
}) => (
	<div>
		<label>{label}</label>
		<div>
			<input {...input} type="date" placeholder={label} max={max} />
			{touched &&
				((error && <span>{error}</span>) ||
					(warning && <span>{warning}</span>))}
		</div>
	</div>
);
