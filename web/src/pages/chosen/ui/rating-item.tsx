import { ExternalLink } from 'lucide-react';

import { Badge } from '#/components/ui/badge';

import type { Rating } from '#/entities/bonds/model';

type RatingItemProps = {
  rating: Rating;
};

export function RatingItem({ rating }: RatingItemProps) {
  return (
    <li className='flex items-start justify-between gap-4 rounded-md border px-3 py-2 hover:bg-muted/50 transition'>
      
      <div className='flex flex-col'>
        <span className='text-sm font-medium leading-tight'>
          {rating.agencyName}
        </span>

        <span className='text-xs text-muted-foreground'>
          {rating.objectName}
        </span>
      </div>

      <div className='flex items-center gap-2'>
        <span className='text-xs text-muted-foreground'>
          {rating.releaseDate.toLocaleDateString()}
        </span>

        <Badge variant='secondary' className='font-semibold'>
          {rating.ratingValue}
        </Badge>

        <a
          href={ rating.releaseUrl }
          target='_blank'
          rel='noreferrer'
          className='text-muted-foreground hover:text-foreground'
        >
          <ExternalLink className='h-4 w-4' />
        </a>
      </div>
    </li>
  );
}