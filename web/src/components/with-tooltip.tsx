import  { TooltipTrigger, TooltipContent, Tooltip } from './ui/tooltip';

import type { ReactNode } from 'react';


type WithTooltipProps = {
  children: ReactNode
  text: string
  side?: 'left' | 'right' | 'bottom' | 'top'
}

export function WithTooltip(props: WithTooltipProps) {
  return (
    <Tooltip>
      <TooltipTrigger asChild>
        {props.children}
      </TooltipTrigger>
      <TooltipContent className='items-center' side={props.side}>
        <p>{props.text}</p>
      </TooltipContent>
    </Tooltip>
  );
}