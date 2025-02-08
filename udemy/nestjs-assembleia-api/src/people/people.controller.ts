import { Body, Controller, Get, Param, Post, Res } from '@nestjs/common';
import { Response } from 'express';
import { PeopleService } from './people.service';
import { Person } from './person';

@Controller('people')
export class PeopleController {
  constructor(private readonly peopleService: PeopleService) {}

  @Get()
  list(@Res() response: Response): Response {
    const list = this.peopleService.list();

    return response.status(200).send(list);
  }

  @Get('/:id')
  findById(@Param('id') id: number, @Res() response: Response): Response {
    const person = this.peopleService.findById(Number(id));

    return response.status(200).send(person);
  }

  @Post()
  save(@Body() person: Person, @Res() response: Response): Response {
    this.peopleService.save(person);

    return response.status(201).send({ message: 'Salvo com sucesso!' });
  }
}
