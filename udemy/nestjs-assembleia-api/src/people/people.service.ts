import { Injectable, NotFoundException } from '@nestjs/common';
import { Person, PersonUpdateRequest } from './person';

@Injectable()
export class PeopleService {
  private readonly people: Person[] = [];

  list(): Person[] {
    return this.people;
  }

  findById(id: number): Person {
    return this.getPersonById(id);
  }

  save(person: Person): void {
    this.people.push(person);
  }

  update(id: number, personUpdateData: PersonUpdateRequest): void {
    const person = this.getPersonById(id);

    Object.assign(person, personUpdateData);
  }

  delete(id: number): void {
    const index = this.people.findIndex((person) => person.id === id);

    if (index === -1) {
      throw new NotFoundException();
    }

    this.people.splice(index, 1);
  }

  private getPersonById(id: number): Person {
    const person = this.people.find((person) => person.id === id);

    if (!person) {
      throw new NotFoundException();
    }

    return person;
  }
}
