import React from "react";
import { MainTile, Tile } from "../components/Filters";

export const Filters = () => {
	const [isKeyDevelop, setIsKeyDevelop] = React.useState<any>([]);
	const [choosedOptions, setChoosedOptions] = React.useState<any[]>([]);
	const switchDevelop = () => {
		setIsKeyDevelop((prev: typeof isKeyDevelop) => {});
	};
	const allfilters = {
		key: ["sub1", "sub2"]
	};
	const choose = () => {
		setChoosedOptions((prev: typeof choosedOptions) => []);
	};

	//gest du Multiple ??? [key, [multiple ? , subs]]
	return (
		<>
			{Object.entries(allfilters).map(([key, subs]) => {
				<>
					<MainTile
						key={key}
						title={key}
						develop={switchDevelop}
						isDevelopped={false}
						choosed={[]}
					/>
					{isKeyDevelop[key as string | number] &&
						subs &&
						subs.map((sub: string) => {
							<Tile
								key={key + "__" + sub}
								title={sub}
								isChoosed={false}
								choose={choose}
							/>;
						})}
				</>;
			})}
			}
		</>
	);
};
