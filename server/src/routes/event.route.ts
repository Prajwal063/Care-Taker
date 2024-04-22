import express from 'express';
import {
    createEvent,
    deleteEvent,
    getAllEvents,
    getEventById,
    registerForEvent,
    unregisterForEvent,
    updateEvent,
} from '../controller/event.controller';
import validate from '../middleware/validate';
import { CreateEventType } from '../types';

const router = express.Router();

router
    .route('/')
    .get(getAllEvents)
    .post(validate(CreateEventType), createEvent);

router
    .route('/:id')
    .get(getEventById)
    .put(validate(CreateEventType), updateEvent)
    .delete(deleteEvent);

router.route('/:id/register').post(registerForEvent);
router.route('/:id/unregister').post(unregisterForEvent);

export default router;
