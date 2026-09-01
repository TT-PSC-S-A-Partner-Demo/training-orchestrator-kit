"""Conventions the three services are supposed to share.

Same function, same contract - normalize() should behave identically everywhere.
Today it does not: orders and billing drifted from the "reject blank loudly" rule
that shipping follows. These tests make the drift visible.

The root cause is not a bug in any one service - it is that the rule was never
written down, so three authors each guessed. Fix it in the spec, then apply once.
"""

import pytest
from services import orders, billing, shipping

MODULES = [pytest.param(orders, id="orders"),
           pytest.param(billing, id="billing"),
           pytest.param(shipping, id="shipping")]


@pytest.mark.parametrize("mod", MODULES)
def test_valid_input_is_trimmed(mod):
    assert mod.normalize("  hello  ") == "hello"


@pytest.mark.parametrize("mod", MODULES)
def test_blank_raises_valueerror(mod):
    with pytest.raises(ValueError):
        mod.normalize("   ")


@pytest.mark.parametrize("mod", MODULES)
def test_none_raises_valueerror(mod):
    with pytest.raises(ValueError):
        mod.normalize(None)
