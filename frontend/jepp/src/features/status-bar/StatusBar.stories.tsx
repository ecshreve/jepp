// StatusBar.stories.ts|tsx
import type { Meta, StoryObj } from '@storybook/react';

import "./StatusBar.css";
import  { StatusBar }  from './StatusBar';

//👇 This default export determines where your story goes in the story list
const meta: Meta<typeof StatusBar> = {
  component: StatusBar,
  tags: ['autodocs'],
  argTypes: { handleClickNewGame: { action: 'clicked' }, handleClickRestart: { action: 'clicked' } },
};

export default meta;
type Story = StoryObj<typeof StatusBar>;

export const Default: Story = {
  args: { gameTitle: "GameID: 4147 -- ShowNum: 234"},
};
