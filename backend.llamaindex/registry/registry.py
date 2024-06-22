from typing import Callable

registry = {}


def register_activity(activity_id: str, func: Callable):
    registry[activity_id] = func


def get_activity(activity_id: str): return registry[activity_id]