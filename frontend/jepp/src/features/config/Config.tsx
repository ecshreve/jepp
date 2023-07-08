import React, { useState } from "react";
import {
	Button,
	ButtonGroup,
	Card,
	Dropdown,
	OverlayTrigger,
	ToggleButton,
	Tooltip,
} from "react-bootstrap";

import "./Config.css";

import {
	setCurrentGameId,
	setGameActive,
} from "./configSlice";
import { useAppDispatch } from "../../app/hooks";
import { YEARS, MONTHS, DEVELOPMENT_GAME_ID } from "../../consts";

export default function Config() {
	const [yearSelection, setYearSelection] = useState(YEARS[0]);
	const [monthSelection, setMonthSelection] = useState(MONTHS[0]);
	const [showNumberSelection, setShowNumberSelection] = useState("Show Number");
	const [selectedGameId, setSelectedGameId] = useState(DEVELOPMENT_GAME_ID);
	const [gameModeSelection, setGameModeSelection] = useState("RANDOM");

	const dispatch = useAppDispatch();
	const renderGameModeSelection = () => {
		const gameModeRadios = [
			{ name: "Random Game", value: "RANDOM" },
			{ name: "Select Game", value: "SELECT" },
		];

		const tooltipDescriptions = [
			"play a random game",
			"choose a specific game",
		];

		return (
			<div>
				<ButtonGroup className="config-gameModeSelection-container">
					{gameModeRadios.map((radio, idx) => (
						<OverlayTrigger
							key={`overlay=${idx}`}
							placement="bottom"
							overlay={
								<Tooltip id={`tooltip-${idx}`}>
									{tooltipDescriptions[idx]}
								</Tooltip>
							}
						>
							<ToggleButton
								id={`toggle-but-${idx}`}
								key={`toggle-but-${idx}`}
								className="config-gameModeSelection-item"
								type="radio"
								variant="secondary"
								name="radio"
								value={radio.value}
								checked={gameModeSelection === radio.value}
								onChange={(e) => setGameModeSelection(e.currentTarget.value)}
							>
								{radio.name}
							</ToggleButton>
						</OverlayTrigger>
					))}
				</ButtonGroup>
			</div>
		);
	};

	const parsed = [{ year: "2020", month: "January", raw: DEVELOPMENT_GAME_ID, showNumber: "#3966"}]
	// const parsedGameIds = useAppSelector((state) => {
		// const gids = state.config.allParsedGameIds;
	// 	return gids.filter(
	// 		(g) => g.year === yearSelection && g.month === monthSelection
	// 	);
	// });

	const renderGameSelection = () => {
		const yearDropdownItems = YEARS.map((y) => {
			return (
				<Dropdown.Item key={y} onClick={() => setYearSelection(y)}>{y}</Dropdown.Item>
			);
		});

		const monthDropdownItems = MONTHS.map((m) => {
			return (
				<Dropdown.Item key={m} onClick={() => setMonthSelection(m)}>{m}</Dropdown.Item>
			);
		});

		const showNumberDropdownItems = parsed.map((g, idx) => {
			return (
				<Dropdown.Item
					key={`showNumberDropdownItem-${idx}`}
					onClick={() => {
						setShowNumberSelection(g.showNumber);
						setSelectedGameId(g.raw);
					}}
				>
					{g.showNumber}
				</Dropdown.Item>
			);
		});

		const gameSelectionDropdownInfo = [
			{ selection: yearSelection, items: yearDropdownItems },
			{ selection: monthSelection, items: monthDropdownItems },
			{ selection: showNumberSelection, items: showNumberDropdownItems },
		];

		const gameSelectionDropdowns = gameSelectionDropdownInfo.map((s, idx) => {
			return (
				<Dropdown id={`dd-${idx}`} key={`dd-${idx}`} as={ButtonGroup} className="config-gameSelection-dropdown">
					<Button
						className="config-gameSelection-dropdown-val"
						variant="success"
					>
						{s.selection}
					</Button>
					<Dropdown.Toggle
						split
						variant="info"
					></Dropdown.Toggle>
					<Dropdown.Menu>{s.items}</Dropdown.Menu>
				</Dropdown>
			);
		});

		return (
			<div>
				<hr />
				<div className="config-gameSelection">{gameSelectionDropdowns}</div>
			</div>
		);
	};

	const renderSaveConfigButton = () => {
		return (
			<div className="config-saveConfig-button">
				<Button
					key="saveConfig"
					variant="primary"
					disabled={selectedGameId === ""}
					onClick={() => {
						console.log(yearSelection, monthSelection, showNumberSelection, selectedGameId)
						dispatch(
							setCurrentGameId(
								gameModeSelection === "SELECT"
									? selectedGameId
									: DEVELOPMENT_GAME_ID
							)
						);
						dispatch(setGameActive(true));
					}}
				>
					Start Game
				</Button>
			</div>
		);
	};

	console.log(gameModeSelection)

	return (
		<>
			<div className="config-container">
				<div className="config-pane">
					<div className="config-pane-header">Welcome to Jeppy!</div>
					<hr />
					<div className="config-pane-content">
						<Card>
							<Card.Body>
								{renderGameModeSelection()}
								{gameModeSelection === "SELECT" && renderGameSelection()}
								<hr />
								{renderSaveConfigButton()}
							</Card.Body>
						</Card>
					</div>
				</div>
			</div>
		</>
	);
}
