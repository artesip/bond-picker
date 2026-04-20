import { Check, Copy } from 'lucide-react';
import { useState } from 'react';

import { Button } from './ui/button';

type CopyButtonProps = {
  value: string
}

export function CopyButton({ value }: CopyButtonProps) {
  const [copied, setCopied] = useState(false);

  const handleCopy = async () => {
    await navigator.clipboard.writeText(value);
    setCopied(true);

    setTimeout(() => setCopied(false), 1500);
  };

  return (
    <Button
      variant='ghost'
      size='icon'
      onClick={ handleCopy }
    >
      {copied ? <Check className='h-3 w-3' /> : <Copy className='h-3 w-3' />}
    </Button>
  );
}