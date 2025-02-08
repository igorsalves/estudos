import { Body, Controller, Get, Post, Res } from '@nestjs/common';
import { Response } from 'express';
import { PeopleService } from './people.service';
import { Person } from './person';

@Controller('people')
export class PeopleController {
  constructor(private readonly peopleService: PeopleService) {}

  @Get()
  list(@Res() response: Response) {
    const list = this.peopleService.list();

    return response.status(200).send(list);
  }

  @Post()
  save(@Body() person: Person, @Res() response: Response) {
    this.peopleService.save(person);

    return response.status(201).send({ message: 'Salvo com sucesso!' });
  }
}
