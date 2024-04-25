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
import { isAuthenticated, validate } from '../middleware';
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

router.route('/:id/register').post(isAuthenticated, registerForEvent);
router.route('/:id/unregister').post(isAuthenticated, unregisterForEvent);

export default router;
