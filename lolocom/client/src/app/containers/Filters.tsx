import React, { Fragment } from "react";
import { MainTile, Tile, FiltersPrev } from "../components/Filters";

export const Filters = () => {
	const [isKeyDevelop, setIsKeyDevelop] = React.useState<any>({});
	const [choosedOptions, setChoosedOptions] = React.useState<any[]>([]);
	const switchDevelop = () => {
		setIsKeyDevelop((prev: typeof isKeyDevelop) => {});
	};
	const allfilters = [
		{ key: "ok", subs: ["sub1", "sub2"], isMultiple: true }
	];
	const choose = (e: any, isMultiple: boolean) => {
		setChoosedOptions((prev: typeof choosedOptions) => []);
	};
	const reInit = () => {};

	return (
		<FiltersPrev>
			{allfilters.map(
				({
					key,
					subs
				}: {
					key: string;
					subs: any[];
					isMultiple: boolean;
				}) => (
					<Fragment key={key}>
						<MainTile
							title={key}
							develop={switchDevelop}
							isDevelopped={false}
							choosed={[]}
						/>
						{isKeyDevelop[key as string | number] &&
							subs &&
							subs.map((sub: string) => (
								<Tile
									key={key + "__" + sub}
									title={sub}
									isChoosed={false}
									choose={choose}
								/>
							))}
						<button onClick={reInit}>ReInit</button>
					</Fragment>
				)
			)}
		</FiltersPrev>
	);
};
