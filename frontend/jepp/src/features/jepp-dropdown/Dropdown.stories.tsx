// Button.stories.ts|tsx
import type { Meta, StoryObj } from '@storybook/react';

import JeppDropdown from './Dropdown';

// More on how to set up stories at: https://storybook.js.org/docs/react/writing-stories/introduction#default-export
const meta: Meta<typeof JeppDropdown> = {
  title: 'UI/Dropdown',
  component: JeppDropdown,
  tags: ['autodocs'],
  argTypes: { setter: { action: 'clicked' } },
};

export default meta;
type Story = StoryObj<typeof JeppDropdown>;

export const Primary: Story = {
  args: {
    options: ['hello', 'world'],
    selection: 'hello',
  },
};
