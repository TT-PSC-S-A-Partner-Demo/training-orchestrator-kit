"""Convention every service must share: normalize() trims valid input and rejects
anything blank. orders, billing and shipping already pass. The flagship task is to
BUILD a fourth service - invoices - that joins them.

invoices does not exist yet, so its rows are red ("not implemented"). Once the agent
team builds it, the rule that matters is whether it matches the convention - and the
convention was never written in the spec, only in the sibling code. That gap is what
the team rewinds on.
"""

import importlib
import pytest

SERVICES = ["orders", "billing", "shipping", "invoices"]


def load(name):
    try:
        return importlib.import_module(f"services.{name}")
    except ModuleNotFoundError:
        return None


@pytest.mark.parametrize("name", SERVICES)
def test_valid_input_is_trimmed(name):
    mod = load(name)
    if mod is None:
        pytest.fail(f"{name} service not implemented yet")
    assert mod.normalize("  hello  ") == "hello"


@pytest.mark.parametrize("name", SERVICES)
def test_blank_is_rejected(name):
    mod = load(name)
    if mod is None:
        pytest.fail(f"{name} service not implemented yet")
    with pytest.raises(ValueError):
        mod.normalize("   ")


@pytest.mark.parametrize("name", SERVICES)
def test_none_is_rejected(name):
    mod = load(name)
    if mod is None:
        pytest.fail(f"{name} service not implemented yet")
    with pytest.raises(ValueError):
        mod.normalize(None)
