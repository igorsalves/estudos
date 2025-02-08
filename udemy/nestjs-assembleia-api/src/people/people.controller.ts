import {
  Body,
  Controller,
  Delete,
  Get,
  Param,
  Post,
  Put,
  Res,
} from '@nestjs/common';
import { Response } from 'express';
import { PeopleService } from './people.service';
import { Person, PersonUpdateRequest } from './person';

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

  @Put('/:id')
  update(
    @Param('id') id: number,
    @Body() personUpdateData: PersonUpdateRequest,
    @Res() response: Response,
  ): Response {
    this.peopleService.update(Number(id), personUpdateData);

    return response.status(200).send({ message: 'Atualizado com sucesso!' });
  }

  @Delete('/:id')
  delete(@Param('id') id: number, @Res() response: Response): Response {
    this.peopleService.delete(Number(id));

    return response.status(204).send();
  }
}
