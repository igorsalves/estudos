import { Controller, Get, Res } from '@nestjs/common';
import { Response } from 'express';
import { PeopleService } from './people.service';

@Controller('people')
export class PeopleController {
  constructor(private readonly peopleService: PeopleService) {}

  @Get()
  list(@Res() response: Response) {
    const list = this.peopleService.list();

    return response.status(200).send(list);
  }
}
